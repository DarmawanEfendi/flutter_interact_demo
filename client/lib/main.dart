import 'package:flutter/material.dart';
import 'package:swagger/api.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Startup Name Generator',
      home: RandomWords(),
      theme: ThemeData(
        primaryColor: Colors.white
      ),
    );
  }
}

class RandomWordsState extends State<RandomWords> {
  var _suggestions = List<ProfileSchema>();
  final _biggerFont = const TextStyle(fontSize: 18.0);

  RandomWordsState() {
    // This widget is the root of your application.
    try {
      var profileApi = new ProfileApi();
      profileApi.getProfiles().then((data) {
        this._suggestions = data.data;
      });
    } catch (e) {
      print("Exception when calling ProfileApi->getProfiles: $e\n");
    }
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Startup Name Generator'),
        actions: <Widget>[
          IconButton(icon: Icon(Icons.list), onPressed: _pushSaved),
        ],
      ),
      body: _buildSuggestions(),
    );
  }

  void _pushSaved() {
    Navigator.of(context).push(
      MaterialPageRoute<void>(
        builder: (BuildContext context) {
          final Iterable<ListTile> tiles = _suggestions.map(
              (ProfileSchema pair) {
                if (pair.isFavourite) {
                  return ListTile(
                    title: Text(pair.name.toUpperCase(), style: _biggerFont,),
                  );
                }
                return null;
              }
          );
          final List<Widget> divided = ListTile.divideTiles(tiles: tiles, context: context).toList();
          return Scaffold(
            appBar: AppBar(title: Text('Saved Suggestions'),),
            body: ListView(children: divided,),
          );
        },
      ),
    );
  }

  Widget _buildSuggestions() {
    return ListView.builder(
      padding: const EdgeInsets.all(16.0),
      itemCount: _suggestions.length,
      itemBuilder: (context, i) {
        if (i.isOdd) return Divider();
        return _buildRow(_suggestions[i], i);
      },
    );
  }

  Widget _buildRow(ProfileSchema pair, int i) {
    final bool alreadySaved = pair.isFavourite;
    return ListTile(
      title: Text(
        pair.name.toUpperCase(),
        style: _biggerFont,
      ),
      subtitle: Text(
        pair.url
      ),
      trailing: Icon(
        alreadySaved ? Icons.favorite: Icons.favorite_border,
        color: alreadySaved ? Colors.red : null,
      ),
      onTap: () {
        try {
          var profileApi = new ProfileApi();
          pair..isFavourite =  !pair.isFavourite;
          profileApi.postProfileFavourite(pair.id, pair).then((data) {
            print("Success: $data");
            setState(() {
              _suggestions[i] = data.data;
            });
          });
        } catch (e) {
          print("Exception when calling ProfileApi->postProfile: $e\n");
        }
      },
    );
  }
}

class RandomWords extends StatefulWidget {
  @override
  State<StatefulWidget> createState() {
    return RandomWordsState();
  }

}