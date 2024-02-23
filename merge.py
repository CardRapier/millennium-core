import os
import sys

def merge_sql_files(output_file):
    input_folder = './db/schemas'  # Replace with the actual path to your SQL files folder
    merged_content_up = []
    merged_content_down = []
    up = True

    for file_name in os.listdir(input_folder):
        if file_name.endswith(".sql"):
            file_path = os.path.join(input_folder, file_name)
            with open(file_path, 'r') as file:
                content = file.read().strip().split('\n')
                for line in content:
                    if line.startswith("-- +goose"):
                        if 'Up' in line:
                            up = True
                        elif 'Down' in line:
                            up = False
                    else:
                        if up:
                            merged_content_up.append(line)
                        else:
                            merged_content_down.append(line)
    # Clear the contents of the output file
    with open(f"./db/migrations/{output_file}.sql", 'w') as clear_output:
        clear_output.write("")

    # Write the merged content to the output file
    with open(f"./db/migrations/{output_file}.sql", 'a') as output:
        output.write('-- +goose Up\n-- +goose StatementBegin\n')
        output.write("\n".join(merged_content_up))
        output.write('-- +goose StatementEnd')
        output.write('\n-- +goose Down\n-- +goose StatementBegin\n')
        output.write("\n".join(merged_content_down))
        output.write('\n-- +goose StatementEnd')

    print(f'Merged content written to {output_file}')

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python merge_sql_files.py <output_file_path>")
    else:
        output_file_path = sys.argv[1]
        merge_sql_files(output_file_path)
