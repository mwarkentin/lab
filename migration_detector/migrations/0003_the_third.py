from django.db import migrations, models

class Migration(migrations.Migration):

    dependencies = [('migrations', '0002_the_second')]

    operations = [
        migrations.DeleteModel('Tribble'),
        migrations.AddField('Author', 'rating', models.IntegerField(default=0)),
    ]
