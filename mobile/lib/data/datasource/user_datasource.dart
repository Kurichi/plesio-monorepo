import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:kiikuten/data/model/user_model.dart';

class UserDataSource {
  final String baseUrl;

  UserDataSource(this.baseUrl);

  Future<UserModel> getUser(String userId) async {
    final response = await http.get(Uri.parse('$baseUrl/users/$userId'));

    if (response.statusCode == 200) {
      return UserModel.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load user');
    }
  }
}
