import 'package:flutter/material.dart';
import 'mqtt_service_2.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  final mqttService = MqttService();

  @override
  void initState() {
    super.initState();
    mqttService.connect();
  }

  @override
  void dispose() {
    mqttService.disconnect();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: Text('Flutter MQTT')),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              ElevatedButton(
                onPressed: () {
                  mqttService.subscribe('test/topic/lutfan');
                },
                child: Text('Subscribe'),
              ),
              ElevatedButton(
                onPressed: () {
                  mqttService.publish('test/topic/lutfan', 'Hello from Flutter');
                },
                child: Text('Publish'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
