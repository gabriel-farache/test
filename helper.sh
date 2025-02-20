#!/bin/bash

ApiHost="http://localhost:8080"

# Get All Widgets
AllWidgets=$(curl -s $ApiHost/api/v1/widgets)
echo $AllWidgets

# Get Widget by ID
ThirdWidget=$(curl -s $ApiHost/api/v1/widgets/176)
echo $ThirdWidget

# Try to get a widget that doesn't exist
curl -s $ApiHost/api/v1/widgets/39289

# Create a new widget
Widget='{
    "name": "Widget 1",
    "description": "This is a widget",
    "count": 10
}'
CreateResponse=$(curl -si -X POST -H "Content-Type: application/json" -d "$Widget" $ApiHost/api/v1/widgets | grep -oP 'Location: \K.*')
NewWidgetUrl="$ApiHost$(echo $CreateResponse)"
NewWidget=$(curl -s $NewWidgetUrl)
echo $NewWidget

# Update a new widget
UpdateWidget='{
  "ID": '"$(echo $NewWidget | jq '.ID')"',
  "name": "Widget 2",
  "description": "This is a widget that has been updated",
  "count": 50
}'
CreateResponse=$(curl -s -X PUT -H "Content-Type: application/json" -d "$UpdateWidget" "$ApiHost/api/v1/widgets/$(echo $UpdateWidget | jq '.ID')" | grep -oP 'Location: \K.*')
UpdatedWidgetUrl="$ApiHost$(echo $CreateResponse)"
UpdatedWidget=$(curl -s $UpdatedWidgetUrl)
echo $UpdatedWidget

# Delete Widget by ID
DeleteResponse=$(curl -s -X DELETE "$ApiHost/api/v1/widgets/$(echo $UpdateWidget | jq '.ID')")
echo $DeleteResponse