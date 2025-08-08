import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text("Flutter App"),
        ),
        body: ListView(
          children: [
            ListTile(
              title: Text("List Tile 1"),
              subtitle: Text('This is a subtitle ksnka skjdnak dkasj dskdjsndka sdkjas '),
              leading: CircleAvatar(
                backgroundImage: NetworkImage(
                  'https://images.unsplash.com/photo-1534528741775-53994a69daeb?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60'
                ),
              ),
              trailing: Text('10:00 AM'),
              dense: true,
            ),
            Divider(),
            ListTile(
              title: Text("List Tile 1"),
              subtitle: Text('This is a subtitle'),
              leading: CircleAvatar(
                backgroundImage: NetworkImage(
                  'https://images.unsplash.com/photo-1534528741775-53994a69daeb?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60'
                ),
              ),
              trailing: Text('10:00 AM'),
            ),
            Divider(),
            ListTile(
              title: Text("List Tile 1"),
              subtitle: Text('This is a subtitle'),
              leading: CircleAvatar(
                backgroundImage: NetworkImage(
                  'https://images.unsplash.com/photo-1534528741775-53994a69daeb?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60'
                ),
              ),
              trailing: Text('10:00 AM'),
            ),
            Divider(),
            ListTile(
              title: Text("List Tile 1"),
              subtitle: Text('This is a subtitle'),
              leading: CircleAvatar(
                backgroundImage: NetworkImage(
                  'https://images.unsplash.com/photo-1534528741775-53994a69daeb?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60'
                ),
              ),
              trailing: Text('10:00 AM'),
            ),
            Divider(),
          ],
        ),
      ),
    );
  }
}