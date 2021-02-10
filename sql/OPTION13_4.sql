BEGIN TRANSACTION;
CREATE TABLE NormEventsData ("serial_number" INTEGER NOT NULL, "event_timestamp" TEXT,"event_date" TEXT,"temp-avg" REAL, "active_pwr-avg"	REAL, "wind_dir-avg" REAL, "available_pwr-avg" REAL, PRIMARY KEY("serial_number","event_timestamp"), FOREIGN KEY ("serial_number") REFERENCES NormSelectorData ("serial_number"));
COMMIT;
BEGIN TRANSACTION;
INSERT INTO NormEventsData ("serial_number", "event_timestamp", "event_date", "temp-avg", "active_pwr-avg", "wind_dir-avg", "available_pwr-avg") SELECT serial_number, event_timestamp, event_date, "temp-avg", "active_pwr-avg", "wind_dir-avg", "available_pwr-avg" FROM EventsData;
COMMIT;