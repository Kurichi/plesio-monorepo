import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:kiikuten/data/model/mission_model.dart';

class MissionDataSource {
  final String baseUrl;

  MissionDataSource(this.baseUrl);

  Future<List<MissionModel>> getMissions() async {
    final response = await http.get(Uri.parse('$baseUrl/missions'));

    if (response.statusCode == 200) {
      Iterable jsonResponse = json.decode(response.body);
      return List<MissionModel>.from(
        jsonResponse.map((model) => MissionModel.fromJson(model)),
      );
    } else {
      throw Exception('Failed to load missions');
    }
  }
}
