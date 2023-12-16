import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:kiikuten/data/model/item_model.dart';
import 'package:kiikuten/domain/entity/mission.dart';

part 'mission_model.freezed.dart';

part 'mission_model.g.dart';

@freezed
class MissionModel with _$MissionModel {
  const MissionModel._();

  const factory MissionModel({
    required String id,
    required String description,
    required String term,
    required String target,
    required int amount,
    required String unit,
    required List<ItemModel> rewards,
  }) = _MissionModel;

  factory MissionModel.fromJson(Map<String, dynamic> json) =>
      _$MissionModelFromJson(json);

  Mission toEntity() {
    return Mission(
      id: id,
      description: description,
      term: term,
      target: target,
      amount: amount,
      unit: unit,
      rewards: rewards.map((e) => e.toEntity()).toList(),
    );
  }
}

@freezed
class UserMissionModel with _$UserMissionModel {
  const UserMissionModel._();

  const factory UserMissionModel({
    required String userId,
    required MissionModel mission,
    required int progress,
    required int deadline,
  }) = _UserMissionModel;

  factory UserMissionModel.fromJson(Map<String, dynamic> json) =>
      _$UserMissionModelFromJson(json);

  UserMission toEntity() {
    return UserMission(
      userId: userId,
      mission: mission.toEntity(),
      progress: progress,
      deadline: deadline,
    );
  }
}
