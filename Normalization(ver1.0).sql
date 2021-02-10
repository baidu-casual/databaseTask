INSERT INTO "Norm_SelectorData" ("serial_number","super_id", "cust_id")
SELECT sd.serial_number, sd.super_id, sd.cust_id FROM SelectorData sd;

INSERT INTO "Norm_Country_Data" ("super_id", "selector_id", "Country", "Model", "Capacity")
SELECT DISTINCT SelectorData.super_id, Events.selector_id, SelectorData.country, SelectorData.model, SelectorData.capacity
FROM Events
INNER JOIN SelectorData ON Events.serial_number=SelectorData.serial_number;


