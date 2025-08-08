import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  List<Tab> myTab = [
    Tab(icon: Icon(Icons.home)),
    Tab(icon: Icon(Icons.search)),
    Tab(icon: Icon(Icons.settings)),
    Tab(icon: Icon(Icons.notifications)),
  ];
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: DefaultTabController(
        // initialIndex: 2,
        length: myTab.length, 
        child: Scaffold(
          appBar: AppBar(
            title: Text('My App'),
            centerTitle: true,
            backgroundColor: Colors.blue,
            bottom: TabBar(
              indicatorColor: Colors.white,
              labelColor: Colors.white,
              unselectedLabelColor: Colors.grey,
              indicatorWeight: 5,
              tabs: myTab,
            ),
          ),
          body: TabBarView(
            children:[
              Center(child: Text(
                'Home', 
                style: TextStyle(fontSize: 48, fontWeight: FontWeight.bold))),
              Center(child: Text('Search', style: TextStyle(fontSize: 48))),
              Center(child: Text('Settings', style: TextStyle(fontSize: 48))),
              Center(child: Text('Notifications', style: TextStyle(fontSize: 48))),
            ]
          )
        ),
      )
    );
  }
}