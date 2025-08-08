import 'package:fl_chart/fl_chart.dart';
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
      home: HomePage(),
    );
  }
}

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Scaffold(
      body: Center(
        child: AspectRatio(
          aspectRatio: 2.0,
          child: LineChart(
            LineChartData(lineBarsData: [
              LineChartBarData(
                  show: true,
                  spots: [
                    FlSpot(0, 0),
                    FlSpot(1, 3),
                    FlSpot(2, 6),
                    FlSpot(3, 2),
                    FlSpot(4, 100),
                    FlSpot(5, 0),
                    FlSpot(6, 100),
                    FlSpot(7, 2),
                  ],
                  // color: Colors.red,
                  gradient: LinearGradient(colors: [
                    Colors.red,
                    Colors.purple,
                    Colors.blueAccent,
                  ], begin: Alignment.bottomCenter, end: Alignment.topCenter),
                  barWidth: 4,
                  isCurved: true,
                  belowBarData: BarAreaData(
                      show: true, color: Colors.blue.withOpacity(0.5)),
                  dotData: const FlDotData(
                      show: true,
                      // checkToShowDot: (spot, barData) {

                      //   return spot;
                      // }
                    ),
                  dashArray: [4, 3, 15],
                  // isStrokeCapRound: true,
                  // isStrokeJoinRound: true
                  preventCurveOverShooting: true
                )
              ]
            ),
          ),
        ),
      ),
    );
  }
}
