import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget{
    @override
    Widget build(BuildContext context) {
        // TODO: implement build
        return MaterialApp(
            home: Scaffold(
                body: Center(
                    child: Text(
                        "Hello World ld skadjak sjldkan sdnasld alksdn", 
                        maxLines: 2,
                        overflow: TextOverflow.ellipsis,
                        textAlign: TextAlign.center,
                        style: TextStyle(
                            backgroundColor: Colors.red,
                            color: Colors.white,
                            fontSize: 30,
                            fontWeight: FontWeight.bold,
                            letterSpacing: 2
                        ),
                    ),
                ),
                appBar: AppBar(
                    backgroundColor: Colors.blue,
                    centerTitle: true,
                    title: Text(
                        "Hello World",
                    ),
                ),
            ),
        );
    }
}