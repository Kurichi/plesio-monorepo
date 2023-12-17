import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:kiikuten/data/model/user_model.dart';

class UserDataSource {
  final String baseUrl;

  UserDataSource(this.baseUrl);

  Future<void> signUp(/* parameters for sign up */) async {
    // Implement sign up logic
    // HTTP POST request to '$baseUrl/api/v1/signup'
  }

  Future<UserModel> getUser(String userId) async {
    final response = await http.get(Uri.parse('$baseUrl/api/v1/user/$userId'));

    if (response.statusCode == 200) {
      return UserModel.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load user');
    }
  }

  Future<void> updateUser(/* parameters for update */) async {
    // Implement update user logic
    // HTTP PUT request to '$baseUrl/api/v1/users/me'
  }
}
