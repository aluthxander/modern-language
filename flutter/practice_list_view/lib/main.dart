import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  List<Widget> myList = [
    Container(
      height: 300,
      width: 300,
      color: Colors.red,
    ),
    Text('Hello World'),
    Container(
      height: 300,
      width: 300,
      color: Colors.green,
    ),
    Container(
      height: 300,
      width: 300,
      color: Colors.blue,
    ),
    Container(
      height: 300,
      width: 300,
      color: Colors.yellow,
    )
  ];

  List<Color> colorList = [
    Colors.red,
    Colors.green,
    Colors.blue,
    Colors.yellow
  ];
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      home: Scaffold(
          appBar: AppBar(
            title: Text('List View'),
          ),
          body: ListView.separated(
            separatorBuilder: (context, index) {
              return Divider(
                height: 20,
              );
            },
            itemCount: 4,
            itemBuilder: (context, index) {
              return Container(
                height: 300,
                width: 300,
                color: colorList[index],
              );
            },
          )),
    );
  }
}
