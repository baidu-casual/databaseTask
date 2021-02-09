#!/bin/bash



git clone https://github.com/wiremoons/csv2sql.git

rm -r sql
rm -r bin

mv csv2sql csv2sqlconv

go build ./csv2sqlconv/csv2sql.go

./csv2sql -f csv/selector_data.csv -t SelectorData

./csv2sql -f csv/events_data.csv -t EventsData

./csv2sql -f csv/events.csv -t Events

rm -r csv2sqlconv

mkdir bin
mv csv2sql bin/
mkdir sql
mv *.sql sql/
