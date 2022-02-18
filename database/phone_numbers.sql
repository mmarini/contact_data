-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS phone_numbers_id_seq;

-- Table Definition
CREATE TABLE "public"."phone_numbers" (
    "id" int4 NOT NULL DEFAULT nextval('phone_numbers_id_seq'::regclass),
    "contact_id" int4 NOT NULL,
    "phone_number" varchar NOT NULL,
    PRIMARY KEY ("id")
);

-- Indexes
CREATE INDEX phone_numbers_by_contact_id ON public.phone_numbers USING btree (contact_id);

