import 'package:flutter/material.dart';

main(){
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          centerTitle: true,
          title: Text("Flutter App"),
        ),
        body: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            Row(
              children: [
                Column(
                  children: [
                    Container(
                      width: 100,
                      height: 100,
                      color: Colors.red,
                    ),
                    Container(
                      width: 200,
                      height: 100,
                      color: Colors.green,
                    ),
                    Container(
                      width: 300,
                      height: 100,
                      color: Colors.blue,
                    ),
                    Container(
                      width: 400,
                      height: 100,
                      color: Colors.yellow,
                    )
                  ]
                )
              ],
            ),
            SizedBox(width: 20),
            Row(
              children: [
                Container(
                  width: 100,
                  height: 100,
                  color: Colors.red,
                ),
                Container(
                  width: 200,
                  height: 100,
                  color: Colors.green,
                ),
                Container(
                  width: 300,
                  height: 100,
                  color: Colors.blue,
                ),
              ],
            ),
            Stack(
              children: [
                Container(
                  width: 20,
                  height: 20,
                  color: Colors.red,
                ),
                Container(
                  width: 40,
                  height: 40,
                  color: Colors.green,
                ),
                Container(
                  width: 60,
                  height: 60,
                  color: Colors.blue,
                ),
              ],
            )
          ]
        ),
      ),
    );
  }
}