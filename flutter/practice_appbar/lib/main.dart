import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          leading: Container(
            color: Colors.amberAccent,
          ),
          leadingWidth: 100,
          title: Text("My APP"),
          centerTitle: true,
          actions: [
            Container(
              width: 35,
              color: Colors.amber,
            ),
            Icon(Icons.search),
          ],
          bottom: PreferredSize(
            preferredSize: Size.fromHeight(200), 
            child: Container(
              width: 50,
              height: 200,
              color: Colors.amber,
            )
          ),
          flexibleSpace: Container(
              width: 50,
              height: 200,
              color: Colors.green,
              child: Text("hallo"),
          ),
        ),
      ),
    );
  }
}