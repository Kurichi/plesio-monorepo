import 'package:kiikuten/data/datasource/mission_datasource.dart';
import 'package:kiikuten/domain/entity/mission.dart';
import 'package:kiikuten/domain/repository/mission_repository.dart';

class MissionRepositoryImpl implements MissionRepository {
  final MissionDataSource _missionDataSource;

  MissionRepositoryImpl(this._missionDataSource);

  @override
  Future<List<Mission>> getMissions() async {
    final missionModels = await _missionDataSource.getMissions();
    return missionModels.map((e) => e.toEntity()).toList();
  }

  @override
  Future<UserMission> getUserMission(String userId) async {
    final userMissionModel = await _missionDataSource.getUserMission(userId);
    return userMissionModel.toEntity();
  }

  @override
  Future<void> updateUserMission(UserMission userMission) async {
    //
  }
}
