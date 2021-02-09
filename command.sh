#!/bin/bash

echo "Removing sql directory..."
rm -r sql
echo "Removing bin directory..."
rm -r bin

git clone https://github.com/wiremoons/csv2sql.git

mv csv2sql csv2sqlconv

echo "Building csv2sql.go..."
go build ./csv2sqlconv/csv2sql.go

./csv2sql -f csv/selector_data.csv -t SelectorData

./csv2sql -f csv/events_data.csv -t EventsData

./csv2sql -f csv/events.csv -t Events

echo "Removing csv2sqlconv directory..."
rm -r -f csv2sqlconv

echo "Creating bin directory..."
mkdir bin
mv csv2sql bin/
echo "Creating sql directory..."
mkdir sql
mv *.sql sql/
