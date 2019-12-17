part of swagger.api;

class MetaSchema {
  /* HTTP response code */
  int status = null;
  
/* Message of response */
  String message = null;
  
  MetaSchema();

  @override
  String toString() {
    return 'MetaSchema[status=$status, message=$message, ]';
  }

  MetaSchema.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    status =
        json['status']
    ;
    message =
        json['message']
    ;
  }

  Map<String, dynamic> toJson() {
    return {
      'status': status,
      'message': message
     };
  }

  static List<MetaSchema> listFromJson(List<dynamic> json) {
    return json == null ? new List<MetaSchema>() : json.map((value) => new MetaSchema.fromJson(value)).toList();
  }

  static Map<String, MetaSchema> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, MetaSchema>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new MetaSchema.fromJson(value));
    }
    return map;
  }
}

