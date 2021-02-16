SELECT eventsdata.serial_number, eventsdata.event_timestamp, selectordata.country
FROM eventsdata
LEFT JOIN selectordata ON selectordata.serial_number = eventsdata.serial_number
ORDER BY selectordata.country;