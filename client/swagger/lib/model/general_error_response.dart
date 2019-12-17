part of swagger.api;

class GeneralErrorResponse {
  
  MetaSchema meta = null;
  
  GeneralErrorResponse();

  @override
  String toString() {
    return 'GeneralErrorResponse[meta=$meta, ]';
  }

  GeneralErrorResponse.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    meta =
      
      
      new MetaSchema.fromJson(json['meta'])
;
  }

  Map<String, dynamic> toJson() {
    return {
      'meta': meta
     };
  }

  static List<GeneralErrorResponse> listFromJson(List<dynamic> json) {
    return json == null ? new List<GeneralErrorResponse>() : json.map((value) => new GeneralErrorResponse.fromJson(value)).toList();
  }

  static Map<String, GeneralErrorResponse> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, GeneralErrorResponse>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new GeneralErrorResponse.fromJson(value));
    }
    return map;
  }
}

