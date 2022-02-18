-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS contacts_id_seq;

-- Table Definition
CREATE TABLE "public"."contacts" (
    "id" int4 NOT NULL DEFAULT nextval('contacts_id_seq'::regclass),
    "full_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    PRIMARY KEY ("id")
);

