SELECT DISTINCT selectordata.serial_number, selectordata.super_id, events.selector_id
FROM selectordata
INNER JOIN events ON events.serial_number = selectordata.serial_number;