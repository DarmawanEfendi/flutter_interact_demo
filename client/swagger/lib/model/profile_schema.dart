part of swagger.api;

class ProfileSchema {
  /* id of profile */
  int id = null;
  
/* name of profile */
  String name = null;
  
/* url of profile */
  String url = null;
  
/* status of favourite */
  bool isFavourite = null;
  
  ProfileSchema();

  @override
  String toString() {
    return 'ProfileSchema[id=$id, name=$name, url=$url, isFavourite=$isFavourite, ]';
  }

  ProfileSchema.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    id =
        json['id']
    ;
    name =
        json['name']
    ;
    url =
        json['url']
    ;
    isFavourite =
        json['is_favourite']
    ;
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'url': url,
      'is_favourite': isFavourite
     };
  }

  static List<ProfileSchema> listFromJson(List<dynamic> json) {
    return json == null ? new List<ProfileSchema>() : json.map((value) => new ProfileSchema.fromJson(value)).toList();
  }

  static Map<String, ProfileSchema> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, ProfileSchema>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new ProfileSchema.fromJson(value));
    }
    return map;
  }
}

