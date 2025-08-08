import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget{
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text("Date Formate"),
        ),
        body: Center(
          child: Text(
            DateFormat.EEEE().format(DateTime.now()),
            style: TextStyle(
              fontSize: 25,
            ),
          ),
        ),
      )
    );
  }
}
