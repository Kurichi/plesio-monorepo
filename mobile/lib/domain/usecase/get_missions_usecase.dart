import 'package:kiikuten/domain/entity/mission.dart';
import 'package:kiikuten/domain/repository/mission_repository.dart';
import 'package:kiikuten/provider/repository/mission_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'get_missions_usecase.g.dart';

@riverpod
GetMissionsUseCase getMissionsUseCase(GetMissionsUseCaseRef ref) {
  final missionRepository = ref.watch(missionRepositoryProvider);
  return GetMissionsUseCase(missionRepository);
}

class GetMissionsUseCase {
  final MissionRepository _missionRepository;

  GetMissionsUseCase(this._missionRepository);

  Future<List<Mission>> execute() async {
    return await _missionRepository.getMissions();
  }
}
