BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "employees" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"password"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "roles" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"borrow_day"	integer,
	"bookroom_hr"	integer,
	"bookcom_hr"	integer,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "roles" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "member_classes" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"discount"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"pin"	text,
	"first_name"	text,
	"last_name"	text,
	"civ"	text,
	"phone"	text,
	"email"	text,
	"password"	text,
	"address"	text,
	"emp_id"	integer,
	"role_id"	integer,
	"role_id"	integer,
	"member_class_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_roles_users" FOREIGN KEY("role_id") REFERENCES "roles"("id"),
	CONSTRAINT "fk_member_classes_users" FOREIGN KEY("member_class_id") REFERENCES "member_classes"("id"),
	CONSTRAINT "fk_employees_users" FOREIGN KEY("emp_id") REFERENCES "employees"("id"),
	CONSTRAINT "fk_roles_users" FOREIGN KEY("role_id") REFERENCES "roles"("id")
);
CREATE INDEX IF NOT EXISTS "idx_employees_deleted_at" ON "employees" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_roles_deleted_at" ON "roles" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_roles_deleted_at" ON "roles" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_member_classes_deleted_at" ON "member_classes" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_email" ON "users" (
	"email"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_civ" ON "users" (
	"civ"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_pin" ON "users" (
	"pin"
);
COMMIT;
