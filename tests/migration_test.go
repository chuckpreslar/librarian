package test

import (
  "testing"
  "strings"
)

import(
  "github.com/chuckpreslar/librarian"
)

func TestCreateTableBasic(t *testing.T) {
  migrator := new(librarian.Migrator)
  users := migrator.CreateTable("users")

  expected := `CREATE TABLE "users" ();`

  if sql, err := users.Sequalize(); nil != err || strings.Trim(sql, "\n") != expected {
    t.Errorf("Returned unexpected error (%v) or sql (%v), wanted %v.", err, sql, expected)
  }
}

func TestCreateTableWithColumns(t *testing.T) {
  migrator := new(librarian.Migrator)
  users := migrator.CreateTable("users").
    String("first_name", librarian.ColumnOptions{ Null: false }).
    String("last_name", librarian.ColumnOptions{ Null: false }).
    String("email", librarian.ColumnOptions{ Null: false }).
    Boolean("banned", librarian.ColumnOptions{ Default: false }).
    Integer("id")

  expected := "CREATE TABLE \"users\" ();\n" +
              "ALTER TABLE \"users\" ADD \"first_name\" character varying(255);\n" +
              "ALTER TABLE \"users\" ADD \"last_name\" character varying(255);\n" +
              "ALTER TABLE \"users\" ADD \"email\" character varying(255);\n" +
              "ALTER TABLE \"users\" ADD \"banned\" boolean;\n" +
              "ALTER TABLE \"users\" ADD \"id\" integer;\n" +
              "ALTER TABLE \"users\" ALTER \"first_name\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"last_name\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"email\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"banned\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"banned\" SET DEFAULT 'false';"

  sql, err := users.Sequalize()

  if nil != err {
    t.Errorf("Returned unexpected error: %v\n", err)
  } else if sql != expected {
    t.Errorf("Returned unexpected SQL:\n%v\n", sql)
    t.Errorf("Expected:\n%v\n", expected)
  }
}

func TestCreateTableWithUnnamedUniqueIndex(t *testing.T) {
  migrator := new(librarian.Migrator)
  users := migrator.CreateTable("users").
    String("first_name", librarian.ColumnOptions{ Null: false }).
    String("last_name", librarian.ColumnOptions{ Null: false }).
    String("email", librarian.ColumnOptions{ Null: false }).
    Boolean("banned", librarian.ColumnOptions{ Default: false }).
    Integer("id").
    AddUniqueIndex("email").
    AddPrimaryKeyIndex("id").
    AddForeignKeyIndex("test_id", "test")

  expected := "CREATE TABLE \"users\" ();\n" +
              "ALTER TABLE \"users\" ADD \"first_name\" character varying(255);\n" +
              "ALTER TABLE \"users\" ADD \"last_name\" character varying(255);\n" +
              "ALTER TABLE \"users\" ADD \"email\" character varying(255);\n" +
              "ALTER TABLE \"users\" ADD \"banned\" boolean;\n" +
              "ALTER TABLE \"users\" ADD \"id\" integer;\n" +
              "ALTER TABLE \"users\" ALTER \"first_name\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"last_name\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"email\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"banned\" SET NOT NULL;\n" +
              "ALTER TABLE \"users\" ALTER \"banned\" SET DEFAULT 'false';\n" +
              "ALTER TABLE \"users\" ADD CONSTRAINT \"users_unique_email\" UNIQUE(\"email\");\n" +
              "ALTER TABLE \"users\" ADD CONSTRAINT \"users_pkey_id\" PRIMARY KEY(\"id\");\n" +
              "ALTER TABLE \"users\" ADD CONSTRAINT \"users_fkey_test_id\" FOREIGN KEY(\"test_id\") REFERENCES \"test\";"

  sql, err := users.Sequalize()

  if nil != err {
    t.Errorf("Returned unexpected error: %v\n", err)
  } else if sql != expected {
    t.Errorf("Returned unexpected SQL:\n%v\n", sql)
    t.Errorf("Expected:\n%v\n", expected)
  }
}


