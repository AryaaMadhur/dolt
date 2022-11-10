#!/usr/bin/env bats
load $BATS_TEST_DIRNAME/helper/common.bash

# Test case for running autogenerated queries against a table

setup() {
  setup_common

  mkdir test
  cd test
  dolt init

  # Load in a complex database
  dolt sql <<SQL
CREATE TABLE \`default_table\` (
  \`pk\` int NOT NULL DEFAULT '2',
  \`col2\` float NOT NULL DEFAULT (LENGTH('hello')) COMMENT 'fsdsdf',
  \`col3\` double NOT NULL DEFAULT (ROUND(-1.58, 0)),
  \`col4\` float DEFAULT (RAND()) COMMENT 'fsdfsd',
  \`col5\` int DEFAULT '4' COMMENT 'dass',
  PRIMARY KEY (\`pk\`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO \`default_table\` (\`pk\`,\`col2\`,\`col3\`,\`col4\`,\`col5\`) VALUES (1,5,-2,0.14581752,2);
INSERT INTO \`default_table\` (\`pk\`,\`col2\`,\`col3\`,\`col4\`,\`col5\`) VALUES (2,5,-2,0.5525121,2);
INSERT INTO \`default_table\` (\`pk\`,\`col2\`,\`col3\`,\`col4\`,\`col5\`) VALUES (3,5,-2,0.9616921,3);
INSERT INTO \`default_table\` (\`pk\`,\`col2\`,\`col3\`,\`col4\`,\`col5\`) VALUES (4,5,-2,0.93345636,4);
CREATE TABLE \`keyed_table\` (
  \`pk\` int NOT NULL,
  \`v2\` binary(1),
  \`v1\` bigint,
  \`v4\` blob,
  \`v3\` bit(1),
  \`v5\` tinyint,
  \`v9\` decimal(10,0),
  \`v8\` datetime,
  \`v7\` char(1),
  \`v6\` date,
  \`v15\` int,
  \`v12\` float,
  \`v10\` double,
  \`v11\` json,
  \`v13\` linestring,
  \`v16\` longtext,
  \`v14\` longblob,
  \`v17\` mediumblob,
  \`v19\` mediumtext,
  \`v29\` tinyint,
  \`v18\` mediumint,
  \`v20\` mediumtext,
  \`v21\` point,
  \`v23\` smallint,
  \`v22\` polygon,
  \`v25\` text,
  \`v24\` smallint,
  \`v26\` time(6),
  \`v27\` timestamp,
  \`v28\` tinyblob,
  \`v30\` tinyint,
  \`v31\` tinytext,
  \`v32\` varchar(255),
  \`v33\` year,
  PRIMARY KEY (\`pk\`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO \`keyed_table\` (\`pk\`,\`v2\`,\`v1\`,\`v4\`,\`v3\`,\`v5\`,\`v9\`,\`v8\`,\`v7\`,\`v6\`,\`v15\`,\`v12\`,\`v10\`,\`v11\`,\`v13\`,\`v16\`,\`v14\`,\`v17\`,\`v19\`,\`v29\`,\`v18\`,\`v20\`,\`v21\`,\`v23\`,\`v22\`,\`v25\`,\`v24\`,\`v26\`,\`v27\`,\`v28\`,\`v30\`,\`v31\`,\`v32\`,\`v33\`) VALUES (1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);
INSERT INTO \`keyed_table\` (\`pk\`,\`v2\`,\`v1\`,\`v4\`,\`v3\`,\`v5\`,\`v9\`,\`v8\`,\`v7\`,\`v6\`,\`v15\`,\`v12\`,\`v10\`,\`v11\`,\`v13\`,\`v16\`,\`v14\`,\`v17\`,\`v19\`,\`v29\`,\`v18\`,\`v20\`,\`v21\`,\`v23\`,\`v22\`,\`v25\`,\`v24\`,\`v26\`,\`v27\`,\`v28\`,\`v30\`,\`v31\`,\`v32\`,\`v33\`) VALUES (2,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);
CREATE TABLE \`keyless\` (
  \`pk\` int,
  \`val\` int,
  KEY \`myidx\` (\`val\`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO \`keyless\` (\`pk\`,\`val\`) VALUES (3,2);
INSERT INTO \`keyless\` (\`pk\`,\`val\`) VALUES (1,1);
INSERT INTO \`keyless\` (\`pk\`,\`val\`) VALUES (1,1);
INSERT INTO \`keyless\` (\`pk\`,\`val\`) VALUES (2,1);
CREATE TABLE \`trigger_table\` (
  \`x\` int NOT NULL,
  PRIMARY KEY (\`x\`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE VIEW myview AS SELECT * FROM \`keyed_table\`;
SQL
}
# executed after each test
teardown() {
    assert_feature_version
    teardown_common
}

@test "tableplus: load all the tables with no errors" {
  run dolt sql -q "SELECT * FROM \`test\`.\`default_table\` ORDER BY \`pk\` LIMIT 300 OFFSET 0;"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT table_rows as count FROM information_schema.TABLES WHERE TABLE_SCHEMA='dsimple'AND TABLE_NAME='default_table';"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT * FROM \`test\`.\`keyed_table\` ORDER BY \`pk\` LIMIT 300 OFFSET 0;"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT table_rows as count FROM information_schema.TABLES WHERE TABLE_SCHEMA='test'AND TABLE_NAME='keyed_table';"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT VIEW_DEFINITION as create_statement FROM INFORMATION_SCHEMA.VIEWS WHERE TABLE_SCHEMA='test'AND TABLE_NAME='myview';"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT * FROM \`test\`.\`myview\` LIMIT 300 OFFSET 0;"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT table_rows as count FROM information_schema.TABLES WHERE TABLE_SCHEMA='test'AND TABLE_NAME='myview';"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT * FROM \`test\`.\`trigger_table\` ORDER BY \`x\` LIMIT 300 OFFSET 0;"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT table_rows as count FROM information_schema.TABLES WHERE TABLE_SCHEMA='test'AND TABLE_NAME='trigger_table';"
  [ "$status" -eq 0 ]

  run dolt sql -q "SELECT COUNT(*) as count FROM \`test\`.\`trigger_table\`;"
  [ "$status" -eq 0 ]
}

@test "tableplus: load data from the keyed table" {
  run dolt sql -r csv -q "SELECT * FROM \`test\`.\`keyed_table\` ORDER BY \`pk\` LIMIT 300 OFFSET 0;"
  [ "$status" -eq 0 ]
  [[ "$output" =~ "pk,v2,v1,v4,v3,v5,v9,v8,v7,v6,v15,v12,v10,v11,v13,v16,v14,v17,v19,v29,v18,v20,v21,v23,v22,v25,v24,v26,v27,v28,v30,v31,v32,v33" ]] || false
  [[ "$output" =~ "1,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,," ]] || false
  [[ "$output" =~ "2,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,," ]] || false

  run dolt sql -r csv -q "SELECT table_rows as count FROM information_schema.TABLES WHERE TABLE_SCHEMA='test' AND TABLE_NAME='keyed_table';"
  [ "$status" -eq 0 ]
  [[ "$output" =~ "count" ]] || false
}

@test "tableplus: open the schema of a table" {
  run dolt sql -r csv -q "SELECT ordinal_position as ordinal_position,column_name as column_name,column_type AS data_type,character_set_name as character_set,collation_name as collation,is_nullable as is_nullable,column_default as column_default,extra as extra,column_name AS foreign_key,column_comment AS comment FROM information_schema.columns WHERE table_schema='test' AND table_name='keyless'"
  [ "$status" -eq 0 ]
  echo $output
  [[  "$output" =~ "ordinal_position,column_name,data_type,character_set,collation,is_nullable,column_default,extra,foreign_key,comment" ]] || false
  [[  "$output" =~ '1,pk,int,,,YES,,"",pk,""' ]] || false
  [[  "$output" =~  '2,val,int,,,YES,,"",val,""' ]] || false

  run dolt sql -r csv -q "SELECT sub_part as index_length,index_name as index_name,index_type AS index_algorithm,CASE non_unique WHEN 0 THEN'TRUE'ELSE'FALSE'END AS is_unique,column_name as column_name FROM information_schema.statistics WHERE table_schema='test' AND table_name='keyless'ORDER BY seq_in_index ASC;"
  [ "$status" -eq 0 ]
  [[  "$output" =~ "index_length,index_name,index_algorithm,is_unique,column_name" ]] || false
  [[  "$output" =~ ",myidx,BTREE,FALSE,val" ]] || false
}

@test "tableplus: views" {
  run dolt sql -r csv -q "SELECT VIEW_DEFINITION as create_statement FROM INFORMATION_SCHEMA.VIEWS WHERE TABLE_SCHEMA='test' AND TABLE_NAME='myview'"
  [ "$status" -eq 0 ]
  [[ "$output" =~ "create_statement" ]] || false
  [[ "$output" =~ "SELECT * FROM \`keyed_table\`" ]] || false
}

@test "tableplus: procedures" {
  run dolt sql -q "CREATE PROCEDURE test.simple_proc1(x DOUBLE, y DOUBLE) SELECT x*y;"
  [ "$status" -eq 0 ]
  run dolt sql -r csv -q "SHOW CREATE procedure simple_proc1"
  [ "$status" -eq 0 ]
  [[ "$output" =~ "CREATE PROCEDURE test.simple_proc1(x DOUBLE, y DOUBLE) SELECT x*y" ]] || false
}

@test "tableplus: foreign key relationships" {
  # Add some foreign keys
  dolt sql <<SQL
    CREATE TABLE colors (
        id INT NOT NULL,
        color VARCHAR(32) NOT NULL,

        PRIMARY KEY (id),
        INDEX color_index(color)
    );
    CREATE TABLE objects (
        id INT NOT NULL,
        name VARCHAR(64) NOT NULL,
        color VARCHAR(32),

        PRIMARY KEY(id),
        FOREIGN KEY (color) REFERENCES colors(color)
    );
    INSERT INTO colors (id,color) VALUES (1,'red'),(2,'green'),(3,'blue'),(4,'purple');
    INSERT INTO objects (id,name,color) VALUES (1,'truck','red'),(2,'ball','green'),(3,'shoe','blue');
SQL

  run dolt sql -r csv -q "SELECT TABLE_NAME,TABLE_SCHEMA,COLUMN_NAME,REFERENCED_TABLE_SCHEMA,REFERENCED_TABLE_NAME,REFERENCED_COLUMN_NAME FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE TABLE_SCHEMA='test' AND REFERENCED_TABLE_NAME IS NOT NULL ORDER BY ORDINAL_POSITION;"
  [ "$status" -eq 0 ]
  [[ "$output" =~ "TABLE_NAME,TABLE_SCHEMA,COLUMN_NAME,REFERENCED_TABLE_SCHEMA,REFERENCED_TABLE_NAME,REFERENCED_COLUMN_NAME" ]] || false
  [[ "$output" =~ "objects,test,color,test,colors,color" ]] || false
}

@test "tableplus: index queries" {
  dolt sql -q "alter table keyless add index (val)"
  run dolt sql -r csv -q "SELECT ordinal_position as ordinal_position,column_name as column_name,column_type AS data_type,character_set_name as character_set,collation_name as collation,is_nullable as is_nullable,column_default as column_default,extra as extra,column_name AS foreign_key,column_comment AS comment FROM information_schema.columns WHERE table_schema='test' AND table_name='keyless';"
  [[ "$output" =~ "ordinal_position,column_name,data_type,character_set,collation,is_nullable,column_default,extra,foreign_key,comment" ]] || false
  [[ "$output" =~ '1,pk,int,,,YES,,"",pk,""' ]] || false
  [[ "$output" =~ '2,val,int,,,YES,,"",val,""' ]] || false
}

@test "tableplus: trigger queries" {
  dolt sql -q "CREATE TRIGGER ins_sum BEFORE INSERT ON test.keyless FOR EACH ROW SET @sum = @sum + NEW.val;"

  run dolt sql -r csv -q "SELECT trigger_name as name,event_manipulation as event,action_timing as timing,action_statement as statement FROM information_schema.triggers WHERE event_object_schema='test' AND event_object_table='keyless';"
  [ "$status" -eq 0 ]
  [[ "$output" =~ 'name,event,timing,statement' ]] || false
  [[ "$output" =~ 'ins_sum,INSERT,BEFORE,SET @sum = @sum + NEW.val' ]] || false
}

@test "tableplus: connection queries" {
  run dolt sql -q "select version()"
  [ "$status" -eq 0 ]
  run dolt sql -q "show session status like 'ssl_version';"
  [ "$status" -eq 0 ]
  run dolt sql -q "SET NAMES utf8mb4;"
  [ "$status" -eq 0 ]
  run dolt sql -q "SHOW VARIABLES;"
  [ "$status" -eq 0 ]
  run dolt sql -q "show databases;"
  [ "$status" -eq 0 ]
  run dolt sql -q "SELECT ROUTINE_SCHEMA as function_schema,ROUTINE_NAME as function_name,ROUTINE_DEFINITION as create_statement,ROUTINE_TYPE as function_type FROM information_schema.routines where ROUTINE_SCHEMA='test';"
  [ "$status" -eq 0 ]
  run dolt sql -q "show full tables from test"
  [ "$status" -eq 0 ]
}

@test "tableplus: information_schema.routines excludes procedure aliases, but works with CALL" {
  run dolt sql -q "SELECT ROUTINE_SCHEMA as function_schema,ROUTINE_NAME as function_name,ROUTINE_DEFINITION as create_statement,ROUTINE_TYPE as function_type FROM information_schema.routines where ROUTINE_SCHEMA='test';"
  [ "$status" -eq 0 ]
  [[ ! "$output" =~ "dadd" ]] || false
  [[ ! "$output" =~ "dadd" ]] || false
  [[ ! "$output" =~ "dbranch" ]] || false
  [[ ! "$output" =~ "dcheckout" ]] || false
  [[ ! "$output" =~ "dclean" ]] || false
  [[ ! "$output" =~ "dcommit" ]] || false
  [[ ! "$output" =~ "dfetch" ]] || false
  [[ ! "$output" =~ "dmerge" ]] || false
  [[ ! "$output" =~ "dpull" ]] || false
  [[ ! "$output" =~ "dpush" ]] || false
  [[ ! "$output" =~ "dreset" ]] || false
  [[ ! "$output" =~ "drevert" ]] || false
  [[ ! "$output" =~ "dverify_constraints" ]] || false
  [[ ! "$output" =~ "dverify_all_constraints" ]] || false

  run dolt sql -q "CALL dbranch('branch1')"
  [ "$status" -eq 0 ]
  [[ "$output" =~ "status" ]] || false
}

@test "tableplus: server with no dbs yet should be able to describe dolt stored procedures" {
    # make directories outside of the existing init'ed dolt repos
    tempDir=$(mktemp -d)
    cd $tempDir
    mkdir repo1
    cd repo1

    # check that without a DB we get descriptive errors
    run dolt sql -q "SHOW CREATE PROCEDURE dolt_clone;"
    [ "$status" -eq 1 ]
    [[ "$output" =~ "no database selected" ]] || false

    run dolt sql -q "SHOW CREATE PROCEDURE dolt_branch;"
    [ "$status" -eq 1 ]
    [[ "$output" =~ "no database selected" ]] || false

    # initialize dolt
    dolt init

    # check that the DB "repo1" exists
    run dolt sql -q "SHOW DATABASES;"
    [ "$status" -eq 0 ]
    [[ "$output" =~ "repo1" ]] || false

    # check that the DB "repo1" is selected
    run dolt sql -q "SELECT DATABASE();"
    [ "$status" -eq 0 ]
    [[ "$output" =~ "repo1" ]] || false
    echo "select database output >>> $output"

    # check that current DB can be used
    run dolt sql -q "SHOW CREATE PROCEDURE dolt_branch;"
    [ "$status" -eq 0 ]

    # check that the qualified DB name can be used
    run dolt sql -q "SHOW CREATE PROCEDURE repo1.dolt_branch;"
    [ "$status" -eq 0 ]

    # check that procedures can be queried from multiple DBs
    run dolt sql -q "CREATE DATABASE repo2;"
    [ "$status" -eq 0 ]
    run dolt sql -q "SHOW CREATE PROCEDURE repo2.dolt_branch;"
    [ "$status" -eq 0 ]
    run dolt sql -q "USE repo1; SHOW CREATE PROCEDURE dolt_branch;"
    [ "$status" -eq 0 ]
}
