def generate_retag_test_cases():
    sources = [
        ("alpine", "Docker Hub (official, no tag)"),
        ("alpine:3.16.2", "Docker Hub (official, with tag)"),
        ("nested/alpine", "Docker Hub (community, no tag)"),
        ("nested/alpine:3.16.2", "Docker Hub (community, with tag)"),
        ("quay.io/alpine", "private registry (flat, no tag)"),
        ("quay.io/alpine:3.16.2", "private registry (flat, with tag)"),
        ("quay.io/nested/alpine", "private registry (nested, no tag)"),
        ("quay.io/nested/alpine:3.16.2", "private registry (nested, with tag)"),
    ]

    destinations = [
        ("example", "Docker Hub (community)"),
        ("gcr.io", "private registry (naked)"),
        ("gcr.io/example", "private registry (flat)"),
        ("gcr.io/example/images", "private registry (nested)"),
    ]

    for destination, destination_description in destinations:
        for source, source_description in sources:
            print(
                f"""        {'{'}
            name: "from {source_description} to {destination_description}",
            source: "{source}",
            target: "{destination}",
            want: "{destination}/{source.split('/', 1)[-1]}",
        {'}'},"""
            )


if __name__ == "__main__":
    generate_retag_test_cases()
