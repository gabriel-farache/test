gcloud auth login --update-adc

$ApiHost = "http://localhost:8080"

# ==================================================================================
# Get All Widgets
# ==================================================================================
$AllWidgets = Invoke-RestMethod -Uri $ApiHost/api/v1/widgets
$AllWidgets

# ==================================================================================
# Get Widget by ID
# ==================================================================================
$ThirdWidget = Invoke-RestMethod -Uri $ApiHost/api/v1/widgets/176
$ThirdWidget

# ==================================================================================
# Try to get a widget that doesn't exist
# ==================================================================================
Invoke-RestMethod -Uri $ApiHost/api/v1/widgets/39289

# ==================================================================================
# Create a new widget
# ==================================================================================
$Widget = @{
    name = "Widget 1"
    description = "This is a widget"
    count = 10
}
$CreateResponse = Invoke-WebRequest -Uri $ApiHost/api/v1/widgets -Method Post -Body ($Widget | ConvertTo-Json -Depth 10 -Compress)
$NewWidgetUrl = "$ApiHost$($CreateResponse.Headers["Location"])"
$NewWidget = Invoke-RestMethod -Uri $NewWidgetUrl
$NewWidget

# ==================================================================================
# Update a new widget
# ==================================================================================
$UpdateWidget = @{
  ID = $NewWidget.ID
  name = "Widget 2"
  description = "This is a widget that has been updated"
  count = 50
}
$CreateResponse = Invoke-WebRequest -Uri "$ApiHost/api/v1/widgets/$($UpdateWidget.ID)" -Method Put -Body ($UpdateWidget | ConvertTo-Json -Depth 10 -Compress)
$UpdatedWidgetUrl = "$ApiHost$($CreateResponse.Headers["Location"])"
$UpdatedWidget = Invoke-RestMethod -Uri $UpdatedWidgetUrl
$UpdatedWidget

# ==================================================================================
# Delete Widget by ID
# ==================================================================================
$DeleteResponse = Invoke-WebRequest -Uri "$ApiHost/api/v1/widgets/$($UpdateWidget.ID)" -Method Delete
$DeleteResponse