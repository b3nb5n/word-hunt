import 'dart:async';

import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(fontFamily: 'Helvetica'),
      home: HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  Timer? _ticker;
  String _letter = 'A';

  @override
  void initState() {
    final duration = Duration(seconds: 1);
    _ticker = Timer.periodic(duration, (timer) {
      final charCode = timer.tick % 26 + 'A'.codeUnitAt(0);
      setState(() {
        _letter = String.fromCharCode(charCode);
      });
    });

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          Container(
            decoration: BoxDecoration(
              image: DecorationImage(
                image: AssetImage('assets/training_base.png'),
                fit: BoxFit.contain,
              ),
            ),
          ),
          Align(
            alignment: Alignment(0, 0.406),
            child: LayoutBuilder(builder: (ctx, constraints) {
              return Container(
                width: constraints.maxWidth * 0.7,
                height: constraints.maxWidth * 0.7,
                child: GridView.count(
                  crossAxisCount: 4,
                  physics: NeverScrollableScrollPhysics(),
                  padding: EdgeInsets.zero,
                  mainAxisSpacing: 9,
                  crossAxisSpacing: 9,
                  children: [
                    for (var i = 0; i < 16; i++)
                      Container(
                        child: Center(
                          child: Text(
                            _letter,
                            style: TextStyle(
                              fontSize: 45,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                        ),
                      ),
                  ],
                ),
              );
            }),
          ),
        ],
      ),
    );
  }
}
