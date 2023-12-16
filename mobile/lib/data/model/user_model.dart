import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:kiikuten/data/model/tree_model.dart';
import 'package:kiikuten/domain/entity/user.dart';

part 'user_model.freezed.dart';

part 'user_model.g.dart';

@freezed
class UserModel with _$UserModel {
  const UserModel._();

  const factory UserModel({
    required String id,
    required String username,
    required String avatarUrl,
    required TreeModel tree,
  }) = _UserModel;

  factory UserModel.fromJson(Map<String, dynamic> json) =>
      _$UserModelFromJson(json);

  User toEntity() {
    return User(
      id: id,
      username: username,
      avatarUrl: avatarUrl,
      tree: tree.toEntity(),
    );
  }
}
