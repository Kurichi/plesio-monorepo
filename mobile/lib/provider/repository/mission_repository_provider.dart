import 'package:kiikuten/data/repository/mission_repository_impl.dart';
import 'package:kiikuten/domain/repository/mission_repository.dart';
import 'package:kiikuten/provider/datasource/mission_datasource_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'mission_repository_provider.g.dart';

@riverpod
MissionRepository missionRepository(MissionRepositoryRef ref) {
  final missionDataSource = ref.watch(missionDataSourceProvider);
  return MissionRepositoryImpl(missionDataSource);
}
