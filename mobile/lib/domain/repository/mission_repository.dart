import 'package:kiikuten/domain/entity/mission.dart';

abstract class MissionRepository {
  Future<List<Mission>> getMissions();
  Future<void> progressMission(String missionId);
}
