SELECT DISTINCT Events.serial_number, Events.selector_id, SelectorData.super_id, SelectorData.country, SelectorData.model, SelectorData.capacity
FROM Events
RIGHT JOIN SelectorData ON Events.serial_number=SelectorData.serial_number
ORDER BY Events.selector_id;