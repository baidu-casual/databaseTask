CREATE TABLE NormEvents ( serial_number INTEGER NOT NULL, event_name TEXT, FOREIGN KEY (serial_number) REFERENCES NormSelectorData (serial_number));
INSERT INTO NormEvents (serial_number,event_name) SELECT serial_number, event_name FROM Events;