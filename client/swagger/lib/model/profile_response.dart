part of swagger.api;

class ProfileResponse {
  
  MetaSchema meta = null;
  

  ProfileSchema data = null;
  
  ProfileResponse();

  @override
  String toString() {
    return 'ProfileResponse[meta=$meta, data=$data, ]';
  }

  ProfileResponse.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    meta =
      
      
      new MetaSchema.fromJson(json['meta'])
;
    data =
      
      
      new ProfileSchema.fromJson(json['data'])
;
  }

  Map<String, dynamic> toJson() {
    return {
      'meta': meta,
      'data': data
     };
  }

  static List<ProfileResponse> listFromJson(List<dynamic> json) {
    return json == null ? new List<ProfileResponse>() : json.map((value) => new ProfileResponse.fromJson(value)).toList();
  }

  static Map<String, ProfileResponse> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, ProfileResponse>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new ProfileResponse.fromJson(value));
    }
    return map;
  }
}

