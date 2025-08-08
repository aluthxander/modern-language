import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  int count = 1;

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Dynamic App'),
        ),
        body: Column(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            Text(
              count.toString(), 
              style: TextStyle(
                  fontSize: 20
                ),
              ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                ElevatedButton(
                  onPressed:(){
                    count++;
                    print(count);
                  },
                  child: Icon(Icons.add),
                ),
                ElevatedButton(
                  onPressed:null,
                  child: Icon(Icons.remove),
                )
              ]
            ),
          ],
        ),
      ),
    );
  }
}
