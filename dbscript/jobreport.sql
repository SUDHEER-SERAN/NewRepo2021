CREATE TABLE "tusers" (
  "empid" int,
  "userid" SERIAL PRIMARY KEY,
  "password" varchar
  "username" varchar
);

CREATE TABLE "temployeedetails" (
  "empid" SERIAL PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar,
  "mobileno" int NOT NULL,
  "address" varchar,
  "role" int
);

CREATE TABLE "tcustomer" (
  "custid" SERIAL PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar,
  "mobileno" int NOT NULL,
  "address" varchar,
  "location" varchar,
  "route" varchar,
  "custtype" int
);

CREATE TABLE "trequest" (
  "requestid" SERIAL PRIMARY KEY,
  "custid" int,
  "requestdate" timestamp,
  "complaint" varchar,
  "typeofservice" int,
  "section" varchar,
  "otheritem" varchar,
  "materialused" varchar,
  "sparepartsused" varchar,
  "workstatus" int,
  "pendingreason" varchar,
  "uddetails" varchar,
  "brokerid" int,
  "itemstakenfromclient" varchar,
  "delivereditems" varchar,
  "deliverydate" timestamp,
  "deliveredby" int,
  "verify" char,
  "verifiedby" int,
  "technicianid" varchar,
  "workdetails" varchar,
  "workstartdate" date,
  "workenddate" date,
  "cancellationdate" date,
  "cancellationreason" varchar,
  "customerapproval" char,
  "overtimereason" varchar,
  "oldreqid" int,
  "careof" int
);

CREATE TABLE "tpayment" (
  "payid" SERIAL PRIMARY KEY,
  "requestid" int,
  "modeofpayment" int,
  "estimationamount" int,
  "agreedamount" int,
  "actualamount" int,
  "advanceamount" int,
  "paymentstatus" int
);

CREATE TABLE "tcharges" (
  "requestid" int,
  "inspectioncharge" int,
  "sparepartsamount" int,
  "additiontoolrent" int,
  "transportcharges" int,
  "latheworkcharge" int,
  "vendorcost" int,
  "brokercharge" int,
  "misc" int,
  "totalamount" int
);

CREATE TABLE "tbrokerdetails" (
  "brokerid" SERIAL PRIMARY KEY,
  "fristname" varchar,
  "lastname" varchar,
  "mobileno" int
);

CREATE TABLE "treference" (
  "refid" SERIAL PRIMARY KEY,
  "reftype" varchar,
  "refdescription" varchar
);

CREATE TABLE "treferencecode" (
  "refcodeid" SERIAL PRIMARY KEY,
  "refid" int,
  "refcode" varchar,
  "refcodedescription" varchar
);

ALTER TABLE "tusers" ADD FOREIGN KEY ("empid") REFERENCES "temployeedetails" ("empid");

ALTER TABLE "temployeedetails" ADD FOREIGN KEY ("role") REFERENCES "treferencecode" ("refcodeid");

ALTER TABLE "tcustomer" ADD FOREIGN KEY ("custtype") REFERENCES "treferencecode" ("refcodeid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("custid") REFERENCES "tcustomer" ("custid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("typeofservice") REFERENCES "treferencecode" ("refcodeid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("workstatus") REFERENCES "treferencecode" ("refcodeid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("brokerid") REFERENCES "tbrokerdetails" ("brokerid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("deliveredby") REFERENCES "temployeedetails" ("empid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("oldreqid") REFERENCES "trequest" ("requestid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("careof") REFERENCES "treferencecode" ("refcodeid");

ALTER TABLE "tpayment" ADD FOREIGN KEY ("requestid") REFERENCES "trequest" ("requestid");

ALTER TABLE "tpayment" ADD FOREIGN KEY ("modeofpayment") REFERENCES "treferencecode" ("refcodeid");

ALTER TABLE "tcharges" ADD FOREIGN KEY ("requestid") REFERENCES "trequest" ("requestid");

ALTER TABLE "treferencecode" ADD FOREIGN KEY ("refid") REFERENCES "treference" ("refid");

ALTER TABLE tcustomer ALTER COLUMN mobileno TYPE bigint;
