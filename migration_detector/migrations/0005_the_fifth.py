from django.db import migrations, models

class Migration(migrations.Migration):

    dependencies = [('migrations', '0004_the_fourth')]

    operations = [
        migrations.DeleteModel('Tribble'),
        migrations.AddField('Author', 'rating', models.IntegerField(default=0)),
    ]
