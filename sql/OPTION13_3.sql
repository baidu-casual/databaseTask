BEGIN TRANSACTION;
CREATE TABLE NormEvents ( "serial_number" INTEGER NOT NULL, "event_name" TEXT, FOREIGN KEY ("serial_number") REFERENCES NormSelectorData ("serial_number"));
COMMIT;
BEGIN TRANSACTION;
INSERT INTO NormEvents ("serial_number","event_name") SELECT serial_number, event_name FROM Events;
COMMIT;