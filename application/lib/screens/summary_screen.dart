import 'package:flutter/material.dart';
import '../services/api_service.dart';
import 'package:mental_health_companion/l10n/app_localizations.dart';

class SummaryScreen extends StatefulWidget {
  const SummaryScreen({super.key});

  @override
  State<SummaryScreen> createState() => _SummaryScreenState();
}

class _SummaryScreenState extends State<SummaryScreen> {
  List<String> summaries = [];

  @override
  void initState() {
    super.initState();
    _loadSummaries();
  }

  Future<void> _loadSummaries() async {
    final data = await ApiService.fetchSummaries();
    setState(() => summaries = data);
  }

  @override
  Widget build(BuildContext context) {
    final loc = AppLocalizations.of(context)!;

    return Scaffold(
      appBar: AppBar(title: Text(loc.summaryTitle)),
      body: summaries.isEmpty
          ? const Center(child: CircularProgressIndicator())
          : ListView.builder(
              itemCount: summaries.length,
              itemBuilder: (context, index) {
                return ListTile(
                  title: Text(loc.sessionX(index + 1)),
                  subtitle: Text(summaries[index]),
                );
              },
            ),
    );
  }
}
