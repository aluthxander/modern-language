import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Flutter Demo'),
        ),
        body: Center(
          child: Container(
            width: 350,
            height: 350,
            color: Colors.amber,
            // child: Image(
            //   fit: BoxFit.cover,
            //   image: AssetImage('images/roti.jpg')
            //   // image: NetworkImage('https://images.unsplash.com/photo-1517248135467-4c7edcad340c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80'),
            // ),
            child: Image.asset("images/roti.jpg"),
          ),
        ),
      ),
    );
  }
}