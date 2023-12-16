import 'package:kiikuten/domain/entity/mission.dart';

abstract class MissionRepository {
  Future<List<Mission>> getMissions();
  Future<UserMission> getUserMission(String userId);
  Future<void> updateUserMission(UserMission userMission);
}
