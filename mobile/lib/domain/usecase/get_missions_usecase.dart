import 'package:kiikuten/domain/entity/mission.dart';
import 'package:kiikuten/domain/repository/mission_repository.dart';

class GetMissionsUseCase {
  final MissionRepository _missionRepository;

  GetMissionsUseCase(this._missionRepository);

  Future<List<Mission>> execute() async {
    return await _missionRepository.getMissions();
  }
}
