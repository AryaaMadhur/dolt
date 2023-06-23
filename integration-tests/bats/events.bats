#!/usr/bin/env bats
load $BATS_TEST_DIRNAME/helper/common.bash
load $BATS_TEST_DIRNAME/helper/query-server-common.bash

make_repo() {
  mkdir "$1"
  cd "$1"
  dolt init
  cd ..
}

setup() {
    skiponwindows "tests are flaky on Windows"
    setup_no_dolt_init
    make_repo repo1
}

teardown() {
    stop_sql_server 1 && sleep 0.5
    rm -rf $BATS_TMPDIR/sql-server-test$$
    teardown_common
}

@test "events: simple insert into table event" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"

    start_sql_server
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert1 ON SCHEDULE EVERY 1 DAY DO INSERT INTO totals (int_col) VALUES (1);"
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "ALTER EVENT insert1 DISABLE; SELECT * FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "DISABLED" ]] || false
}

@test "events: disabling recurring event should not be dropped" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"

    start_sql_server
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert1 ON SCHEDULE EVERY 1 DAY DO INSERT INTO totals (int_col) VALUES (1);"
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "ALTER EVENT insert1 DISABLE; SELECT * FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "DISABLED" ]] || false
}

@test "events: disabling current_timestamp one time event after execution" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"

    start_sql_server
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert9 ON SCHEDULE AT CURRENT_TIMESTAMP DO INSERT INTO totals (int_col) VALUES (9);"
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM totals;"
    [ $status -eq 0 ]
    [[ $output =~ "| 1        |" ]] || false

    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "| 0        |" ]] || false

    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert8 ON SCHEDULE AT CURRENT_TIMESTAMP ON COMPLETION PRESERVE DO INSERT INTO totals (int_col) VALUES (8);"
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM totals;"
    [ $status -eq 0 ]
    [[ $output =~ "| 2        |" ]] || false

    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SHOW CREATE EVENT insert8;"
    [ $status -eq 0 ]
    [[ $output =~ "ON COMPLETION PRESERVE DISABLE" ]] || false

    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "ALTER EVENT insert8 ON COMPLETION NOT PRESERVE; SELECT COUNT(*) FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "| 1        |" ]] || false

    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "ALTER EVENT insert8 ENABLE; SELECT COUNT(*) FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "| 0        |" ]] || false
}

@test "events: disabling future one time event after execution from the scheduler" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"

    start_sql_server
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert9 ON SCHEDULE AT CURRENT_TIMESTAMP + INTERVAL 3 SECOND ON COMPLETION PRESERVE DO INSERT INTO totals (int_col) VALUES (9); SHOW CREATE EVENT insert9;"
    [ $status -eq 0 ]
    [[ $output =~ "ON COMPLETION PRESERVE ENABLE" ]] || false

    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT SLEEP(4);"
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM totals;"
    [ $status -eq 0 ]
    [[ $output =~ "| 1        |" ]] || false

    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "| 1        |" ]] || false

    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SHOW CREATE EVENT insert9;"
    [ $status -eq 0 ]
    [[ $output =~ "ON COMPLETION PRESERVE DISABLE" ]] || false
}

@test "events: recurring event with STARTS and ENDS defined" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"

    start_sql_server
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert1 ON SCHEDULE EVERY 2 SECOND STARTS CURRENT_TIMESTAMP + INTERVAL 2 SECOND ENDS CURRENT_TIMESTAMP + INTERVAL 5 SECOND DO INSERT INTO totals (int_col) VALUES (1);"
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT SLEEP(7);"
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM totals;"
    [ $status -eq 0 ]
    [[ $output =~ "| 2        |" ]] || false

    # should be dropped
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "| 0        |" ]] || false
}

@test "events: recurring event with ENDS defined" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"

    start_sql_server
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert1 ON SCHEDULE EVERY 2 SECOND ENDS CURRENT_TIMESTAMP + INTERVAL 3 SECOND ON COMPLETION PRESERVE DO INSERT INTO totals (int_col) VALUES (1); SELECT SLEEP(5);"
    sleep 2
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM totals;"
    [ $status -eq 0 ]
    [[ $output =~ "| 2        |" ]] || false

    # should be disabled
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT * FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "DISABLED" ]] || false
}

@test "events: checking out a branch should disable all events leaving the working set dirty" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"
    dolt add .
    dolt commit -m "create table"

    start_sql_server
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert1 ON SCHEDULE EVERY 2 SECOND ENDS CURRENT_TIMESTAMP + INTERVAL 3 SECOND ON COMPLETION PRESERVE DO INSERT INTO totals (int_col) VALUES (1); SELECT SLEEP(5);"
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CALL DOLT_COMMIT('-am','commit with an event')"

    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CALL DOLT_CHECKOUT('-b','newbranch')"
    # should be disabled
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT * FROM information_schema.events;"
    [ $status -eq 0 ]
    [[ $output =~ "DISABLED" ]] || false
}

@test "events: events on default branch still runs after switching to non default branch" {
    cd repo1
    dolt sql -q "CREATE TABLE totals (id int PRIMARY KEY AUTO_INCREMENT, int_col int)"
    dolt add .
    dolt commit -m "create table"

    start_sql_server
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CREATE EVENT insert1 ON SCHEDULE EVERY 5 SECOND ENDS CURRENT_TIMESTAMP + INTERVAL 1 MINUTE DO INSERT INTO totals (int_col) VALUES (1);"
    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CALL DOLT_COMMIT('-am','commit with an event')"

    dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CALL DOLT_CHECKOUT('-b','newbranch')"
    # should be disabled
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "SELECT COUNT(*) FROM totals;"
    [ $status -eq 0 ]
    [[ $output =~ "| 1        |" ]] || false

    sleep 5
    # should not be 1
    run dolt sql-client -P $PORT -u dolt --use-db 'repo1' -q "CALL DOLT_CHECKOUT('main'); SELECT COUNT(*) FROM totals;"
    [ $status -eq 0 ]
    [[ ! $output =~ "| 1        |" ]] || false
}
