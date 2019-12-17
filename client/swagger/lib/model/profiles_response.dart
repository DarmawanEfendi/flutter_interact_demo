part of swagger.api;

class ProfilesResponse {
  
  MetaSchema meta = null;
  

  List<ProfileSchema> data = [];
  
  ProfilesResponse();

  @override
  String toString() {
    return 'ProfilesResponse[meta=$meta, data=$data, ]';
  }

  ProfilesResponse.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    meta =
      
      
      new MetaSchema.fromJson(json['meta'])
;
    data =
      ProfileSchema.listFromJson(json['data'])
;
  }

  Map<String, dynamic> toJson() {
    return {
      'meta': meta,
      'data': data
     };
  }

  static List<ProfilesResponse> listFromJson(List<dynamic> json) {
    return json == null ? new List<ProfilesResponse>() : json.map((value) => new ProfilesResponse.fromJson(value)).toList();
  }

  static Map<String, ProfilesResponse> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, ProfilesResponse>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new ProfilesResponse.fromJson(value));
    }
    return map;
  }
}

