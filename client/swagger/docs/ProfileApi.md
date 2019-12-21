# swagger.api.ProfileApi

## Load the API package
```dart
import 'package:swagger/api.dart';
```

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteProfileFavourite**](ProfileApi.md#deleteProfileFavourite) | **DELETE** /profiles/{id}/favourite | delete profile from favourite list
[**getProfiles**](ProfileApi.md#getProfiles) | **GET** /profiles | Find all profiles
[**postProfileFavourite**](ProfileApi.md#postProfileFavourite) | **POST** /profiles/{id}/favourite | add profile to favourite list


# **deleteProfileFavourite**
> ProfileResponse deleteProfileFavourite(id)

delete profile from favourite list



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new ProfileApi();
var id = 56; // int | ID of profile

try { 
    var result = api_instance.deleteProfileFavourite(id);
    print(result);
} catch (e) {
    print("Exception when calling ProfileApi->deleteProfileFavourite: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of profile | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getProfiles**
> ProfilesResponse getProfiles(onlyFavourite)

Find all profiles

Returns array of profile

### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new ProfileApi();
var onlyFavourite = true; // bool | query for filter only favourite

try { 
    var result = api_instance.getProfiles(onlyFavourite);
    print(result);
} catch (e) {
    print("Exception when calling ProfileApi->getProfiles: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyFavourite** | **bool**| query for filter only favourite | [optional] 

### Return type

[**ProfilesResponse**](ProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **postProfileFavourite**
> ProfileResponse postProfileFavourite(id, body)

add profile to favourite list



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new ProfileApi();
var id = 56; // int | ID of profile
var body = new ProfileSchema(); // ProfileSchema | Profile object that needs to be added

try { 
    var result = api_instance.postProfileFavourite(id, body);
    print(result);
} catch (e) {
    print("Exception when calling ProfileApi->postProfileFavourite: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of profile | 
 **body** | [**ProfileSchema**](ProfileSchema.md)| Profile object that needs to be added | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

