import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:kiikuten/data/datasource/mission_datasource.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'mission_datasource_provider.g.dart';

@riverpod
MissionDataSource missionDataSource(MissionDataSourceRef ref) {
  return MissionDataSource(dotenv.env['BASE_URL'] ?? '');
}
