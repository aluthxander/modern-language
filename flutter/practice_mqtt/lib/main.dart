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
  final ValueNotifier<String> messageNotifier = ValueNotifier<String>('');

  @override
  void initState() {
    super.initState();
    mqttService.connect(onMessageReceived: (message) {
      messageNotifier.value = message;
    });
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
              Text('Received Message:', style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
              ValueListenableBuilder<String>(
                valueListenable: messageNotifier,
                builder: (context, value, child) {
                  return Text(value, style: TextStyle(fontSize: 16, color: Colors.blue));
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
