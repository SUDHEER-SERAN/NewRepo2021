CREATE TABLE "tusers" (
  "empid" int,
  "userid" SERIAL PRIMARY KEY,
  "password" varcharacter
);

CREATE TABLE "temployeedetails" (
  "empid" SERIAL PRIMARY KEY,
  "firstname" varchar,
  "lastname" varchar,
  "mobileno" int,
  "address" varchar,
  "role" int
);

CREATE TABLE "tcustomer" (
  "custid" SERIAL PRIMARY KEY,
  "firstname" varchar,
  "lastname" varchar,
  "mobileno" int,
  "address" varchar,
  "location" varchar,
  "route" varchar,
  "type" int
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
  "technicianid" varchar2,
  "workdetails" varchar2,
  "workstartdate" datetime,
  "workenddate" datetime,
  "cancellationdate" datetime,
  "cancellationreason" varchar2,
  "customerapproval" char,
  "overtimereason" varchar2,
  "oldreqid" int[]
);

CREATE TABLE "tpayment" (
  "payid" SERIAL PRIMARY KEY,
  "requestid" int,
  "modeofpayment" id,
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

CREATE TABLE "tservicingitem" (
  "serviceid" SERIAL PRIMARY KEY,
  "itemname" varchar
);

CREATE TABLE "tcustomertype" (
  "id" SERIAL PRIMARY KEY,
  "type" varchar
);

CREATE TABLE "tworkstatus" (
  "id" SERIAL PRIMARY KEY,
  "status" varchar
);

CREATE TABLE "tpaymentmode" (
  "id" SERIAL PRIMARY KEY,
  "payment" varchar
);

CREATE TABLE "tbrokerdetails" (
  "brokerid" SERIAL PRIMARY KEY,
  "fristname" varchar,
  "lastname" varchar,
  "mobileno" int
);

ALTER TABLE "tusers" ADD FOREIGN KEY ("empid") REFERENCES "temployeedetails" ("empid");

ALTER TABLE "tcustomer" ADD FOREIGN KEY ("type") REFERENCES "tcustomertype" ("type");

ALTER TABLE "trequest" ADD FOREIGN KEY ("custid") REFERENCES "tcustomer" ("custid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("typeofservice") REFERENCES "tservicingitem" ("serviceid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("workstatus") REFERENCES "tworkstatus" ("id");

ALTER TABLE "trequest" ADD FOREIGN KEY ("brokerid") REFERENCES "tbrokerdetails" ("brokerid");

ALTER TABLE "trequest" ADD FOREIGN KEY ("deliveredby") REFERENCES "temployeedetails" ("empid");

ALTER TABLE "tpayment" ADD FOREIGN KEY ("requestid") REFERENCES "trequest" ("requestid");

ALTER TABLE "tpayment" ADD FOREIGN KEY ("modeofpayment") REFERENCES "tpaymentmode" ("id");

ALTER TABLE "tcharges" ADD FOREIGN KEY ("requestid") REFERENCES "trequest" ("requestid");
