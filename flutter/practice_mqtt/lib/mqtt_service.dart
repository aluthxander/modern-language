import 'package:mqtt_client/mqtt_client.dart';
import 'package:mqtt_client/mqtt_server_client.dart';

class MqttService {
  final MqttServerClient client =
      MqttServerClient('test.mosquitto.org', 'flutter_client');

  Function(String)? onMessageReceived;

  Future<void> connect({Function(String)? onMessageReceived}) async {
    this.onMessageReceived = onMessageReceived; // Simpan callback
    
    client.logging(on: true);
    client.keepAlivePeriod = 20;

    final connMessage = MqttConnectMessage()
        .withClientIdentifier('flutter_client')
        .startClean()
        .withWillQos(MqttQos.atLeastOnce);
    
    client.connectionMessage = connMessage;

    try {
      await client.connect();
      client.subscribe('test/topic/lutfan', MqttQos.atMostOnce);
      
      // Mendengarkan pesan masuk
      client.updates!.listen((List<MqttReceivedMessage<MqttMessage>>? messages) {
        final MqttPublishMessage recMessage = messages![0].payload as MqttPublishMessage;
        final payload = MqttPublishPayload.bytesToStringAsString(recMessage.payload.message);

        if (onMessageReceived != null) {
          onMessageReceived!(payload); // Panggil callback untuk UI
        }
      });

    } catch (e) {
      print('Connection failed: $e');
      client.disconnect();
    }
  }

  void disconnect() {
    client.disconnect();
  }
}
