SELECT events.selector_id, eventsdata.*
FROM events
RIGHT JOIN eventsdata ON eventsdata.serial_number = events.serial_number
ORDER BY events.selector_id