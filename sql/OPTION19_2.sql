SELECT distinct serial_number,AVG(`active_pwr-avg`) AS actual_generation, AVG(`available_pwr-avg`) AS expected_generation
FROM EventsData
WHERE actual_generation<expected_generation
group by serial_number;
