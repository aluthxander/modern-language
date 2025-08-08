import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  final List<Map<String, dynamic>> myList = [
    {
      "Name": "Flutter",
      "Age": 20,
      "favColor": [
        "Black",
        "Amber",
        "Red",
        "BLue",
        "white",
        "green",
        "Black",
        "Amber",
        "Red"
      ]
    },
    {
      "Name": "Dart",
      "Age": 20,
      "favColor": ["White", "Amber", "Green"]
    },
    {
      "Name": "Kotlin",
      "Age": 20,
      "favColor": [
        "Black",
        "Amber",
        "Red",
        "BLue",
        "white",
        "green",
        "Black",
        "Amber",
        "Red"
      ]
    },
    {
      "Name": "Flutter",
      "Age": 20,
      "favColor": [
        "Black",
        "Amber",
        "Red",
        "BLue",
        "white",
        "green",
        "Black",
        "Amber",
        "Red"
      ]
    },
    {
      "Name": "Flutter",
      "Age": 20,
      "favColor": [
        "Black",
        "Amber",
        "Red",
        "BLue",
        "white",
        "green",
        "Black",
        "Amber",
        "Red"
      ]
    },
    {
      "Name": "Flutter",
      "Age": 20,
      "favColor": [
        "Black",
        "Amber",
        "Red",
        "BLue",
        "white",
        "green",
        "Black",
        "Amber",
        "Red"
      ]
    },
  ];

  const MyApp({super.key});
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Flutter Demo Home Page'),
        ),
        body: ListView(
          children: myList.map((data) {
            List myFaforColor = data["favColor"];
            return Card(
              margin: EdgeInsets.all(20),
              child: Container(
                padding: EdgeInsets.all(8),
                child : Column(
                  children: [
                    Row(
                      children: [
                        CircleAvatar(),
                        SizedBox(width: 14,),
                        Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(data["Name"]),
                            Text(data["Age"].toString()),
                          ],
                        )
                      ],
                    ),
                    SingleChildScrollView(
                      scrollDirection: Axis.horizontal,
                      child: Row(
                        children: myFaforColor.map((color) {
                          return Container(
                            color: Colors.amber,
                            margin:
                                EdgeInsets.symmetric(vertical: 15, horizontal: 8),
                            padding: EdgeInsets.all(10),
                            child: Text(color),
                          );
                        }).toList(),
                      ),
                    )
                  ],
                ),
              )
            );
          }).toList(),
        ),
      ),
    );
  }
}
